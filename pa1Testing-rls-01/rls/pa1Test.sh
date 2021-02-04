go build pa1.go
echo "build of pa1.go complete"
# do 2 processes
./pa1 c2-rr.in c2-rr.stu
diff c2-rr.stu c2-rr.base
echo " 2 processes - rr complete"
./pa1 c2-sjf.in c2-sjf.stu
diff c2-sjf.stu c2-sjf.base
echo " 2 processes - sjf complete"
./pa1 c2-fcfs.in c2-fcfs.stu
diff c2-fcfs.stu c2-fcfs.base
echo " 2 processes - fcfs complete"
# now do five processes
./pa1 c5-rr.in c5-rr.stu
diff c5-rr.stu c5-rr.base
echo " 5 processes - rr complete"
./pa1 c5-sjf.in c5-sjf.stu
diff c5-sjf.stu c5-sjf.base
echo " 5 processes - sjf complete"
./pa1 c5-fcfs.in c5-fcfs.stu
diff c5-fcfs.stu c5-fcfs.base
echo " 5 processes - fcfs complete"
# now do 10 processes
./pa1 c10-rr.in c10-rr.stu
diff c10-rr.stu c10-rr.base
echo "10 processes - rr complete"
./pa1 c10-sjf.in c10-sjf.stu
diff c10-sjf.stu c10-sjf.base
echo "10 processes - sjf complete"
./pa1 c10-fcfs.in c10-fcfs.stu
diff c10-fcfs.stu c10-fcfs.base
echo "10 processes - fcfs complete"
