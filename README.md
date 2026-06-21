# trec

trec is a time recorder with labels and results.

This is a small tool I built to record test times and results at work—hence the name 'trec' (test/time recorder). 
I hope it helps others who do similar tasks.

## Usage

```
$ trec test1
Recording... 1m15s
Input memo: OK
Recorded.

$ trec test2
Recording... 10m44s
Input memo: NG
Recorded.

$ trec -l
test1 OK 0:01:15
test2 NG 0:10:44
```

## Mode

Default mode is a recording, i.e. could omit mode from command line argument.

| Mode | Mode (Alt.) | Manual |
| ---- | ----------- | ------ |
| -r   | --recoding  | [manual/recording.txt](manual/recording.txt) |
| -l   | --lookup    | [manual/lookup.txt](manual/lookup.txt) |
| -e   | --edit      | [manual/edit.txt](manual/edit.txt) |
| -d   | --delete    | [manual/delete.txt](manual/delete.txt) |
