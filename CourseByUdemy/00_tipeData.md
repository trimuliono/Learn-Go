# Tipe Data Pada Go-Lang

## Tipe Data Number
Ada dua tipe data number;
- Integer
- Floating Point

### Tipe data Integer

| Tipe Data | Niali Minimum        | Nilai Maximum       |
| --------- | -------------------- | ------------------- |
| int8      | -128                 | 127                 |
| int16     | -32768               | 32767               |
| int32     | -2147483648          | 2147483647          |
| int64     | -9223372036854775808 | 9223372036854775807 |

--- 

**Unsigned Integers** (often called "uints") are just like integers (whole numbers) but have the property that they don't have a + or - sign associated with them. Thus they are always non-negative (zero or positive). We use uint's when we know the value we are counting will always be non-negative.

| Tipe Data | Niali Minimum        | Nilai Maximum       |
| --------- | -------------------- | ------------------- |
| uint8      | 0                     | 255                 |
| uint16     | 0                     | 65535               |
| uint32     | 0                     | 4294967295          |
| uint64     | 0                     | 18446744073709551615 |

---

### Tipe data floating point
Tipe data floating point adalah tipe data yang bisa diisi dengan nilai desimal (nilai yang ada komanya).

| Tipe Data   | Nilai Minimum         | Nilai Maksimum        |
|-------------|-----------------------|-----------------------|
| float32     | 1.18×10⁻³⁸           | 3.4×10³⁸             |
| float64     | 2.23×10⁻³⁰⁸          | 1.80×10³⁰⁸          |
| complex64   | complex numbers with float32 real and imaginary parts.                   |
| complex128  | complex numbers with float64 real and imaginary parts.                   |

