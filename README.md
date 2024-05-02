# Directed graph cycle detection using DFS

## Usage

We use this algorithm to detect cycles in a directed graph.

## Concept

We use a recursive depth-first search (DFS) algorithm to traverse the graph. We keep track of the nodes we have visited in the current path. If we encounter a node that is already in the current path, then we have found a cycle.

## Complexity

Time: O(V + E) \
Space: O(V)

V - number of vertices aka nodes \
E - number of edges aka connections

## Historical context

This algorithm is used in compilers to detect circular dependencies in code. Was first proposed by Robert Tarjan in 1972. \
\
Robert Tarjan is a computer scientist who is known for his work in graph theory and data structures. He has made many contributions to the field of computer science, including the development of algorithms for solving problems related to graphs and networks.

## Class diagram

In my implementation i have graph in string to []string map format, so i dont have any class diagram.
