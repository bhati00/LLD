# Systematic Plan to tackle any LLD Problem  

- ### Understanding the nature of the problem  

  Main goal in this step to determing the nature of the problem  
 we need to ask some questions to ourselves like 

1. is it just a plain object oriented design problem
2. is it process oriented/ workflow like order system/ notification system. 
3. or is it a state driven (like atm machine etc.)
4. or is it requires some real time simulation  

- ### Clarify requirements (functional + non functional )  
  - first we need to understand our core features
  - then move to the understand the other requirements like concurrency, scalibilty etc.   

- ### Identify the Core entities and responsibility
  once we understood our problem in depth, now we need to focus on implementation part.  
  Think in terms of :   
  1. Nouns - they are likely to become classes
  2. Verbs - they become our functions 
  3. Responsibilities can be mapped using SOLID principles (especially SRP).  

- ### Define Relationships and Behaviors  
  Use UML Class Diagram, here our target to make design modular and flexible 

- ### Desginging core classes (interfaces)     
   Have clear responsibilities (Single Responsibility Principle).  
Expose minimal necessary methods (Encapsulation).  
Use interfaces where future flexibility is needed.  







