// Stack Implementation
// Found on developer.apple.com
// Slightly modified to meet my needs

struct Stack<Elemet> {
  var items = [Element]()

  mutating func push(_ item: Element) {
    items.append(item)
  }

  mutating func pop() -> Element {
    return items.removeLast()
  }
}
