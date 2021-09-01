# Go_CrackMe
Little Reversing CrackMe written in GO<br/>

## Info:
I created little reversing crackme challenge written in GO language for educational purpose.<br/>
All compiled programs are not stripped of Symbols so also IDA like 7.5 should recover function names.<br/>
Because I know that IDA Free has only 64bit decompiler available and I also know that
some of you guys are Linux lover - I compiled the crackme program for 2 OS and 2 architectures:<br/>
crackme_go2.exe - Windows OS - 32bit: [[Sample fopen.exe]](https://github.com/Dump-GUY/Malware-analysis-and-Reverse-engineering/blob/main/Tracing%20C%20function%20fopen/fopen.exe)<br/>
crackme_go2x64.exe - Windows OS - 64bit: [[Sample fopen.exe]](https://github.com/Dump-GUY/Malware-analysis-and-Reverse-engineering/blob/main/Tracing%20C%20function%20fopen/fopen.exe)<br/>
crackme_go2_linux - Linux OS - 32bit: [[Sample fopen.exe]](https://github.com/Dump-GUY/Malware-analysis-and-Reverse-engineering/blob/main/Tracing%20C%20function%20fopen/fopen.exe)<br/>
crackme_go2_linux64 - Linux OS - 64bit: [[Sample fopen.exe]](https://github.com/Dump-GUY/Malware-analysis-and-Reverse-engineering/blob/main/Tracing%20C%20function%20fopen/fopen.exe)<br/>

You can nicely learn new things, compare 32bit vs 64bit, learn how strings are referenced in compiled GO programs, how arguments are pushed/moved to function etc...<br/>
Very interesting is for example moving function argumnents to stack (not pushing).<br/>

Only for educational purpose and if you feel really lost, I also included GO src code:[[Sample fopen.exe]](https://github.com/Dump-GUY/Malware-analysis-and-Reverse-engineering/blob/main/Tracing%20C%20function%20fopen/fopen.exe)<br/>


Good Luck and Happy Reversing :)<br/>
