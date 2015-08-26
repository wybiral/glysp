(class Person

    (method __init__ (self name)
        (set self.name name))

    (method greet (self)
        (print (+ "Hello " self.name "!"))))


(set person (Person "Bob"))

(person.greet)
