## Problem: Build a Zoo Management System
You are required to implement a simplified Zoo Management System using Object-Oriented Programming principles in Go. Your system should be able to manage various animals, their behaviors, and zoo operations.

🔧 *Requirements*  
<br>

### Encapsulation:

Each animal must have private fields like name, age, and species. Provide methods to safely update and retrieve these details.

Prevent direct access/modification of these fields from outside the package.
### Abstraction:

Design interfaces to abstract animal behavior, e.g., Animal interface with methods like Speak(), Eat(), and Move().

### Inheritance (via Composition):

Create a base struct BaseAnimal that holds common behavior or attributes.

Reuse this struct in different specific animals like Lion, Elephant, Parrot, etc., using composition.

### Polymorphism:

Store multiple animals in a slice of Animal interface type.

Iterate and call common methods (Speak, Eat, etc.) using dynamic dispatch (interface polymorphism).

 ### Functionality to Implement
Add new animals to the zoo (Lion, Elephant, Parrot, etc.)

Display all animals and their behavior using a loop over the interface slice.

Change an animal’s name and age using encapsulated methods.

Demonstrate different behaviors for different animals using polymorphism.

https://www.linkedin.com/posts/prince-bhati-3b6437175_golang-lowleveldesign-systemdesign-activity-7342156975423614976-ZK9A?utm_source=share&utm_medium=member_desktop&rcm=ACoAACmIZHoBbcKAaYyvCdFmvjmT6JT5hLUHqCw

