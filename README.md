# Go_CrackMe
Little Reversing CrackMe written in GO<br/>

## Info:
I created little reversing crackme challenge written in GO language for educational purpose.<br/>
All compiled programs are not stripped of Symbols so also IDA like 7.5 should recover function names.<br/>
<br/>
Because I know that IDA Free has only 64bit decompiler available and I also know that
some of you guys are Linux lover - I compiled the crackme program for 2 OS and 2 architectures:<br/>

crackme_go2.exe - Windows OS - 32bit: [[crackme_go2.exe]](https://github.com/Dump-GUY/Go_CrackMe/blob/main/Binary/crackme_go2.exe)<br/>
crackme_go2x64.exe - Windows OS - 64bit: [[crackme_go2x64.exe]](https://github.com/Dump-GUY/Go_CrackMe/blob/main/Binary/crackme_go2x64.exe)<br/>
crackme_go2_linux - Linux OS - 32bit: [[crackme_go2_linux]](https://github.com/Dump-GUY/Go_CrackMe/blob/main/Binary/crackme_go2_linux)<br/>
crackme_go2_linux64 - Linux OS - 64bit: [[crackme_go2_linux64]](https://github.com/Dump-GUY/Go_CrackMe/blob/main/Binary/crackme_go2_linux64)<br/>

You can nicely learn new things, compare 32bit vs 64bit, learn how strings are referenced in compiled GO programs, how arguments are pushed/moved to function etc...<br/>
Very interesting is for example moving function argumnents to stack (not pushing).<br/>

Only for educational purpose and if you feel really lost, I also included GO src code: [[crackme_go2.go]](https://github.com/Dump-GUY/Go_CrackMe/blob/main/SRC/crackme_go2.go)<br/>
Good Luck and Happy Reversing :)<br/>

## Recommended tools:
[[IDA Pro/Free]](https://hex-rays.com/ida-free/)<br/>
