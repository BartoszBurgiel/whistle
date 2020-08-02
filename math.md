# The formula for the encryption algorithm



## Variables

##### Let $str$ be the string of the length $n_{str}$ be the string that will be encrypted.

##### Let $key$ be the key of the length $n_{key}$that the string $str$ will be encripted with.

##### Let $s$ be the sum of all characters' values as in $s = \sum_i^nstr[i]$

##### Let the byte $b$ be equal to the value of a byte of the string $str$ with the index $i$

##### Let the byte $k$ be equal to the value of a byte of the string $key$ with the index $i$

##### Let the integer $i$ be the index of the byte that is currenty being encoded.

##### Let the integer $val$ be the base for the encryption of each byte of the string $str$ being equal to $val = b * n_{key}$



## Formula

$$
b_{encrypted} = \Bigg( \sum_{i=0}^{n_{key}-1} \frac{2^{k}}{sum+i_{key}+j} +i_{key}\Bigg) + sum*n_{key} 
$$



## Example

```
str = "Hello world"
key = "270d6b61-be35-4ce1-989f-0d13372f1a17"

enc = "8d07354aa0c030-8cf7c3de604790-8ce855d3989424-8cd8eb292d9090-8cc983de03a3ce-8cba1ff0ffb06e-8caabf610715f1-8c9b622cffade2-8c8c0853cfcd9e-8c7cb1d45e4590-8c6d5ead9260c8"
```


