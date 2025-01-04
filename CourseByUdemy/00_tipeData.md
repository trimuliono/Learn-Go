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

<table>
    <tr>
        <th>Tipe Data</th>
        <th>Nilai Minimum</th>
        <th>Nilai Maksimum</th>
    </tr>
    <tr>
        <td>float32</td>
        <td>1.18×10⁻³⁸</td>
        <td>3.4×10³⁸</td>
    </tr>
    <tr>
        <td>float64</td>
        <td>2.23×10⁻³⁰⁸</td>
        <td>1.80×10³⁰⁸</td>
    </tr>
    <tr>
        <td>complex64</td>
        <td colspan="2">complex numbers with float32 real and imaginary parts.</td>
    </tr>
    <tr>
        <td>complex128</td>
        <td colspan="2">complex numbers with float64 real and imaginary parts.</td>
    </tr>
</table>

---

### Alias

| Tipe Data | Alias untuk   |
|-----------|---------------|
| byte      | uint8         |
| rune      | int32         |
| int       | Minimal int32 |
| uint      | Minimal uint32|

`Minimal int32/uint32` disini maksudnya nilai tergantung arsitektur komputer yg digunakkan apakah 32bit atau 64bit