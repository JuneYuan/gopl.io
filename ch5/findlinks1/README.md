## Why/ How `visit()` fulfills extracting the links recursively?

Mathematical induction may give some help when comprehending 
recursive code. I had [a post](https://juneyuan.github.io/blog/post/recursion-and-induction/) about that.

When it comes to this `visit()` function, we can roughly prove its correctness this way:

+ When there is only one node in the HTML document tree, `visit()` makes it with no doubt;
+ Suppose `visit()` fulfills its duty when there are `K` nodes in the HTML document tree, and we need to prove that  it also deals with `K + 1` nodes correctly. This is not hard:
	+ Think of the `(K + 1)`th node. It may come as a `FirstChild`  or `NextSibling` of someone node. In either case, it will be processed by the `for` loop. The `visit()` function in essence is doing a traversal, kind of like the way DFS does.

## Exercise 5.1

How to change the traversal implementation from *for loop* to *recursion*?

Take a closer look at the data structure of `Node`, provided by *golang.org/x/net/html*. And an example will be much helpful.

![](https://i.postimg.cc/qqvgMzgs/IMG-2068.jpg)

There are two sorts of *Linked List* interweaved here. One starts from any node and goes along its **FirstChild**, then **FirstChild of FirstChild**... (e.g. A -> B -> E, or C -> F). The other type of linked list starts from any node and goes along its **NextSibling** then **NextSibling of NextSibling**... (e.g. B -> C -> D, or F -> G).

Such a interweaved linked list can be converted into a binary tree easily: just take all **FirstChild** nodes as **Left Child** nodes, and **NextSibling** nodes as **Right Child** nodes. Then the DFS traversal code is natural to write.

![convert interweaved linked list to binary tree](https://i.postimg.cc/15NHHT7G/IMG-2069.jpg)

### Test

Can implement the `dfsVisit()` in pre-order/ in-order/ post-order, and check the returned `links`.
