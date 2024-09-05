## Middle on disk Queue

---
---

### Introduction

---

Is it possible to contains near one **1.000** queued strings with different length in RAM? Of course It's easy!!!
But what if we start to contain near **15.000.000** queued strings in RAM? Hmmm.... It will takes much RAM and issue with not enough memory can appear!!!
What if we can save middle (most) part of queue on a disk? It looks tasty! And that I can say about Middle ground Queue (MoDQ). Let's look on MoDQ like on the snake!
Current complex type of data containing can be splited to 4 parts:
> 1. Head.
> 2. Body.
> 3. Tail.
> 4. Nerve system (NS)


