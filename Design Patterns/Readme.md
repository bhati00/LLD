# Systematic Plan to tackle any LLD Problem  

 

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







