(func numbers ()
    (set ch (chan))
    (go
        (set x 0)
        (while true
            (ch (+ "Message from channel: " (str x)))
            (set x (+ x 1))))
    ch)

(set n (numbers))

(for (i 10)
    (print (n)))


