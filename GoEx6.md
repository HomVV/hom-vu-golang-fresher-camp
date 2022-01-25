"# hom-vu-golang-fresher-camp" 
In this course, we shouldn't use foreign key for increasing performance

Some disadvantages of foreign key:
  - It hurts performance of insert, update and delete operations because before those tasks database needs to check if it doesn't violate data integrity.
  - It is difficult to store legacy data from the older databases and sources which may not have been so strict on data quality and integrity
  - Full table reload
