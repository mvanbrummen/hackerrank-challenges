#include <iostream>
#include <stdio.h>
#include <stdlib.h>

using namespace std;

struct Node {
  int data;
  Node *next;
};

struct LinkedList {
  Node *head;
};

void insertNode(struct LinkedList *list, Node *node)
{
  node->next = list->head;
  list->head = node;
}

void printList(struct LinkedList *list)
{
  Node *node = list->head;
  while (node) {
    cout << node << ": Node is " << node->data << " with next " << node->next << endl;
    node = node->next;
  }
}

struct Node *insertTail(struct Node *&head, int data)
{
  Node *newNode = new struct Node;
  newNode->data = data;
  newNode->next = NULL;
  if (head == NULL) {
    head = newNode;
  } else {
    Node *node = head;
     while (node->next) node = node->next;
     node->next = newNode;
 }
  return head;
}

int main()
{
    cout << "Inserting tail on empty list" << endl;
    struct LinkedList *el = new LinkedList;
    insertTail(el->head, 123);
    printList(el);
    delete(el);

    cout << "\nInserting first" << endl;
    struct LinkedList *ll = new LinkedList;
    insertNode(ll, new Node{data: 1});
    insertNode(ll, new Node{data: 2});
    insertNode(ll, new Node{data: 3});
    insertNode(ll, new Node{data: 4});
    printList(ll);

    cout << "\nInserting tail" << endl;
    insertTail(ll->head, 69);
    printList(ll);
    delete(ll);
    return 0;
}
