# Java

## Language convention

## tips

### Avoid ConcurentModificationException

```java
List<String> toRemove = new ArrayList<>();
for (String str : myArrayList) {
    if (someCondition) {
        toRemove.add(str);
    }
}
myArrayList.removeAll(toRemove);

```