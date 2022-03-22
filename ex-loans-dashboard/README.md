# Loan Dashboard

We're a successful peer to peer lending platform giving loans to businesses.
Every day we have multiple loan requests coming in, and would like to see a
summary dashboard with what we managed to originate over the years.

## Input Format

The input to this exercise is as follows:

```
Perky Blenders Coffee, 56101, 100000, B, 8%
Laird Hatters LTD, 14190, 25000, C, 15%, 75000, D, 20%
```

The **input consists of rows** that are **separated by line-feeds/new-lines (`\n`)**.

**Each row has fields** that are **separated by commas** and **each row can vary in length**.

A row consists of the following fields:

1. The business name.
2. The main business nature code.
3. A repeating set of triplets with the loan amount, risk band, and interest
   rate.

For example, the following row is broken down as shown:

```
Laird Hatters LTD, 14190,     25000,        C,        15%,      75000,       D,        20%
|---------------||-------||-----------||---------||--------||-----------||---------||--------|
      NAME         CODE    LOAN AMOUNT  RISK BAND  INTEREST  LOAN AMOUNT  RISK BAND  INTEREST
                          |__________Triplet 1_____________||__________Triplet 2_____________|
```

## Business Natures

* 14190 - Manufacture of other wearing apparel and accessories not elsewhere
  classified
* 47990 - Other retail sale not in stores, stalls or markets
* 56101 - Licensed restaurants
* 75000 - Veterinary activities

## Exercises

We would like you to **perform the following tasks in-order**, however you are encouraged to **think ahead**; i.e would your solution to the initial output task be generic enough to suit the enhanced output task?

### Initial Output

We want to transform this into a standard result that shows:

* The **total** amount we have loaned.
* The business name with the average loan amount we have issued them.

For simplicity, print these out on the screen.

You are allowed to feed the input into your solution however you prefer, e.g.:
- a variable containing the data
- a file
- from stdin
- a test harness

### Enhanced Output

Using a bigger sample:

```
Perky Blenders Coffee, 56101, 100000.10, B, 8%
Laird Hatters, 14190, 25000, C, 15%, 75000, D, 20%
Bobbin Bicycles, 47990, 100000, C, 15%, 50000, B, 7.5%, 80000, A, 3.2%
Arapina Bakery, 56101, 25000, A+, 1.8%
Canine Creche, 75000, 60000, B, 8%
```

Add the following calculations to your result:

* The percentage of money lent per each business nature.
* The Total interest earned.
* The number of loans per each risk band.

### Validation

If there is a problem with the format of the results file then **all good entries
should result in output** and the **error should be displayed with the problem
explained in non-technical language** that an loan analyst might be able to
understand and report back.

**Remember** the input fields are **only separated by commas**; think on different ways the input
can be malformed or formatted.

## Loan Prediction

Suppose we have a loans file from the previous year in the same format. Discuss,
but do not solve in code, how you would you implement a way to predict next
year's result.
