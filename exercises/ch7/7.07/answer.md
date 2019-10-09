### Exercise 7.7: Explain why the help message contains &deg;C when the default value of 20.0 does not.

Because the `value` type for the `CelsiusFlag` is a Celsius struct, utilizing it's `String` method for the usage call.

As an example, I've added a `fahrenheitFlag` to exercise 7.06, using the `celsiusFlag` code structure, to illustrate how this new default uses the &deg;F notation in the help message.


```bash
~$ go build -o ./tempflag ./exercises/ch7/7.06/tempflag2/main.go
~$ ./tempflag
20.0000°C
68.0000°F #default value
~$ ./tempflag -temp 100C -tempf 212F
100.0000°C
212.0000°F
~$ ./tempflag -temp 100N
invalid value "100N" for flag -temp: invalid temperature "100N"
Usage of ~/tempflag:
  -temp value
        the temperature in C (default 20.0000°C)
  -tempf value
        the temperature in F (default 68.0000°F)
exit status 2
~$
```

