### Объектно ориентированные анализ и проектирование (ООАП)

**ООАП** - это методология по проектированию и разработки программного обеспечения основанная на представлении его как совокупности взаимодействующих между собой объектов.

**Абстрактный тип данных (АТД)** - это неявное определение некоторого типа данных в проектируемой системе, которое формально задаёт множество объектов и набор допустимых операций над ними.  

**Класс - это реализация АТД**.

Правильное проектирование системы, в первую очередь, предполагает разработку спецификаций АТД, которые впоследствии будут описывать все наши классы, используемые в проекте.

### Спецификация АТД

- Добавляем **предусловия и постусловия**
- Придерживаемся **принципа достаточной полноты АТД**
- Все методы делим на **конструкторы, запросы и команды**
- По возможности **избегаем явной обработкой исключений в коде**

### АТД популярных структур данных

- [Linked List](abstract_data_types/linked_list/contract.go)
- [Two Way List](abstract_data_types/two_way_list/contract.go)
- [Bloom Filter](abstract_data_types/bloom_filter/contract.go)
- [Bounded Stack](abstract_data_types/bounded_stack/contract.go)
- [Queue](abstract_data_types/queue/contract.go)
- [Deque](abstract_data_types/deque/contract.go)
- [Dyn Array](abstract_data_types/dyn_array/contract.go)
- [Hash Table](abstract_data_types/hash_table/contract.go)
- [Power Set](abstract_data_types/power_set/contract.go)
- [Native Dictionary](abstract_data_types/native_dictionary/contract.go)