/*
"I William Askew (wi357066) afﬁrm that this program is entirely my own work and
 that I have neither developed my code together with any another person, nor copied
 any code from any other person, nor permitted my code to be copied or otherwise
 used by any other person, nor have I copied, modiﬁed,or otherwise used programs
 created by others. I acknowledge that any violation of the above terms will be
 treated as academic dishonesty.”
*/

package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
  "log"
)

type process struct {
  name string
  arrival, burst, wait, turnaround int
  running, arrived, inQueue bool
}

func main() {

  if len(os.Args) < 3 {
    fmt.Println("Missing parameter. Please provide input file name or save file name.")
    return
  }
  file, err := os.Open(os.Args[1])

  if err != nil {
    fmt.Println("Can't read file: ", os.Args[1])
    return
  }
  defer file.Close();

  readInputFile(file);
}

func readInputFile(input *os.File) {
  var p []process

  count := 0
  time := 0
  alg := ""
  q := 0
  scanner := bufio.NewScanner(input)
  var s [][]string
  for scanner.Scan() {
    var x = strings.Fields(scanner.Text())
    s = append(s, x)
   }
   if err := scanner.Err(); err != nil {
     log.Fatal(err)
   }

   for i := 0; i < len(s); i++ {

     if s[i][0] == "end" {
       break
     } else if s[i][0] == "processcount" {
       c, err := strconv.Atoi(s[i][1])

       if err != nil {
         fmt.Println("error on string to int conversion")
       }
       count = c
     } else if s[i][0] == "runfor"{
       t, err := strconv.Atoi(s[i][1])

       if err != nil {
         fmt.Println("error on string to int conversion")
       }
       time = t
     } else if s[i][0] == "use" {
       alg = s[i][1]
     } else if s[i][1] == "name" {
       a, err := strconv.Atoi(s[i][4])

       if err != nil {
         fmt.Println("error on string to int conversion")
       }
       b, err := strconv.Atoi(s[i][6])

       if err != nil {
         fmt.Println("error on string to int conversion")
       }

       proc := process {
         name: s[i][2],
         arrival: a,
         burst: b,
         turnaround: 0,
         wait: 0,
         running: false,
         arrived: false,
         inQueue: false,
       }
       p = append(p, proc)
     } else if s[i][0] == "quantum" {
       quantum, err := strconv.Atoi(s[i][1])
       if err != nil {
         fmt.Println("error on string to int conversion")
       }
       q = quantum
     }
   }

  helper(alg, time, p, q, count)

  return
}

func helper(alg string , time int , p []process, q int, count int) {

  if (alg == "fcfs") {
    fcfs(time, p, count)
  } else if (alg == "rr") {
    rr(time, p, q, count)
  } else if (alg == "sjf") {
    sjf(time, p, count)
  }

  return
}

func fcfs (time int, p []process, count int) {
 f, err := os.Create(os.Args[2])

 if err != nil {
   fmt.Println(err)
   return
 }
 defer f.Close()

 fmt.Fprintf(f, "%3d processes\n", count)
 fmt.Fprintln(f, "Using First-Come First-Served")
 var queue []process
 var running process
 for i := 0; i < time; i++ {
   for j := 0; j < count; j++ {
     if p[j].arrival == i {
       queue = append(queue, p[j])
     }
   }
 }

running = queue[0]
j := 0
for i := 0; i <= time; i++ {

  for n := 0; n < count; n++ {
    if queue[n].arrival == i {
      fmt.Fprintf(f, "Time %3d : %s arrived\n", i, queue[n].name)
    }
    if queue[n].arrival < i && queue[n].running == false {
      queue[n].wait++
    }
  }

  if running.burst == 0 {
    fmt.Fprintf(f, "Time %3d : %s finished\n", i, running.name)
    queue[j] = running
    j++
    if j < count {
      running = queue[j]
    }
  }
  if running.arrival <= i && running.running == false {
    fmt.Fprintf(f, "Time %3d : %s selected (burst %3d)\n", i, running.name, running.burst)
    running.running = true
  }
  if i == time {
    fmt.Fprintf(f, "Finished at time %3d\n\n", i)
  } else if (running.running == false && running.burst != 0)  || j >= count {
    fmt.Fprintf(f, "Time %3d : Idle\n", i)
  }
  if running.running {
    running.turnaround++
    running.burst--
  }
}

for i := 0; i < count; i++ {
  for j := 0; j < count; j++ {
    if queue[i].name == p[j].name {
      p[j] = queue[i]
      p[j].turnaround = p[j].turnaround + p[j].wait
    }
  }
 }

for i := 0; i < count; i++ {
  fmt.Fprintf(f, "%s wait %3d turnaround %3d\n", p[i].name, p[i].wait, p[i].turnaround)
}

return
}

func sjf (time int, p []process, count int) {

  f, err := os.Create(os.Args[2])

  if err != nil {
    fmt.Println(err)
    return
  }
  defer f.Close()


  fmt.Fprintf(f, "%3d processes\n", count)
  fmt.Fprintf(f, "Using preemptive Shortest Job First\n")
  var queue []process
  var running process

  finished := 0

  for i := 0; i <= time; i++ {

    for n := 0; n < count; n++ {
      if p[n].arrival == i && !p[n].inQueue {
        p[n].arrived = true
        fmt.Fprintf(f, "Time %3d : %s arrived\n", i, p[n].name)
        if p[n].burst <= running.burst {
          //select the process and put running process in back of queue
          running.running = false
          queue = append(queue, running)

          running = p[n]
          fmt.Fprintf(f, "Time %3d : %s selected (burst %3d)\n", i, running.name, running.burst)
          running.running = true
        } else if len(queue) == 0 {
          queue = append(queue, p[n])
          p[n].inQueue = true
        } else {
          for x := 0; x < len(queue); x++ {
            //place into queue according to burst
            if p[n].burst < queue[x].burst && !p[n].inQueue {
              p[n].inQueue = true
              queue = insert(queue, p[n], x)
            }
            if x == len(queue)-1 && !p[n].inQueue {
              p[n].inQueue = true
              queue = append(queue, p[n])
            }
          }
        }
      }
    }

    if running.name == ""  || !running.running{
      if len(queue) != 0 {
        running = queue[0]
        queue = remove(queue, 0)

        fmt.Fprintf(f, "Time %3d : %s selected (burst %3d)\n", i, running.name, running.burst)
        running.running = true
      }
    }

    if running.burst == 0 && running.running{
      fmt.Fprintf(f, "Time %3d : %s finished\n", i, running.name)

      finished++

      for n := 0; n < count; n++ {
        if running.name == p[n].name {
          p[n] = running
        }
      }

      if finished != count {
        if len(queue) > 0 {
          running = queue[0]
          queue = remove(queue, 0)

          fmt.Fprintf(f, "Time %3d : %s selected (burst %3d)\n", i, running.name, running.burst)
          running.running = true
        } else {
          running.running = false //no process in queue however processes are expected
        }
      } else {
        running.running = false //no process in queue and no processes expected
      }
    }
    if i == time {
      fmt.Fprintf(f, "Finished at time %3d\n\n", i)
    } else if !running.running || finished == count {
      fmt.Fprintf(f, "Time %3d : Idle\n", i)
    }

    for n := 0; n < len(queue); n++ {
        queue[n].wait++
    }

    if running.running {
      running.burst--
      running.turnaround++
    }
  }//bracket for whole for loop

  for n := 0; n < count; n++ {
   p[n].turnaround = p[n].turnaround + p[n].wait
   fmt.Fprintf(f, "%s wait %3d turnaround %3d\n", p[n].name, p[n].wait, p[n].turnaround)
  }

 return
}

func rr (time int, p []process, q int, count int) {
  f, err := os.Create(os.Args[2])

  if err != nil {
    fmt.Println(err)
    return
  }
  defer f.Close()

  fmt.Fprintf(f, "%3d processes\n", count)
  fmt.Fprintln(f, "Using Round-Robin")
  fmt.Fprintf(f, "Quantum %3d\n\n", q)

  var queue []process
  var running process
  x := 0
  finished := 0

  for i := 0; i <= time; i++ {

    for n := 0; n < count; n++ {
      if p[n].arrival == i && !p[n].inQueue {
        p[n].arrived = true
        p[n].inQueue = true
        fmt.Fprintf(f, "Time %3d : %s arrived\n", i, p[n].name)
        queue = append(queue, p[n])
      }
    }

    if running.name == "" {
      if len(queue) != 0 {
        running = queue[0]
        queue = remove(queue, 0)
        running.running = true
      }
    }

    if x%q == 0 && running.burst != 0 && running.running{
      running.running = false

      queue = append(queue, running)

      if len(queue) > 0 {
        running = queue[0]
        queue = remove(queue, 0)
      }

      fmt.Fprintf(f, "Time %3d : %s selected (burst %3d)\n", i, running.name, running.burst)
      x = 0
      running.running = true
    }
    if running.burst == 0 && running.running{
      fmt.Fprintf(f, "Time %3d : %s finished\n", i, running.name)

      finished++

      for n := 0; n < count; n++ {
        if running.name == p[n].name {
          p[n] = running
        }
      }

      if finished != count {
        if len(queue) > 0 {
          running = queue[0]
          queue = remove(queue, 0)

        fmt.Fprintf(f, "Time %3d : %s selected (burst %3d)\n", i, running.name, running.burst)
        x = 0
        running.running = true
      } else {
        running.running = false //no process in queue however processes are expected
       }
      } else {
        running.running = false //no process in queue and are not processes are expected
      }
    }
    if i == time {
      fmt.Fprintf(f, "Finished at time %3d\n\n", i)
    } else if !running.running || finished == count {
      fmt.Fprintf(f, "Time %3d : Idle\n", i)
    }

    for n := 0; n < len(queue); n++ {
        queue[n].wait++
    }
    if running.running {
      running.burst--
      running.turnaround++
      x++
    }
  }

  for n := 0; n < count; n++ {
   p[n].turnaround = p[n].turnaround + p[n].wait
   fmt.Fprintf(f, "%s wait %3d turnaround %3d\n", p[n].name, p[n].wait, p[n].turnaround)
  }

 return
}

func remove(slice []process, s int) []process {
    return append(slice[:s], slice[s+1:]...)
}

func insert(s []process, proc process, i int) []process{
  var temp process
  s = append(s, temp)
  copy(s[i+1:], s[i:])
  s[i] = proc

  return s
}
