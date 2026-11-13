# learning-go-idiomatic-approach-to-real-world-go-programming

Learning the idiomatic approach to real-world Go programming.

Bu repo "Learning Go" kitobi bo'yicha chapter-chapter olingan konspekt va mashqlardan iborat.
Har bir chapter papkasida mavzu bo'yicha `exercise.go` kodlari va `text.txt` qayd daftari bor.
Quyida shu qaydlarning jamlanmasi keltirilgan.

## Mundarija

- [Chapter 1 — Go toolchain va Makefile](#chapter-1--go-toolchain-va-makefile)
- [Chapter 2 — Literal'lar, tiplar, type conversion va const](#chapter-2--literallar-tiplar-type-conversion-va-const)
- [Chapter 3 — Composite tiplar: array, slice, string, map, struct](#chapter-3--composite-tiplar-array-slice-string-map-struct)
- [Chapter 4 — Blok va shadowing](#chapter-4--blok-va-shadowing)
- [Chapter 5 — Funksiyalar va defer](#chapter-5--funksiyalar-va-defer)
- [Chapter 6 — Pointer, stack/heap va escape analysis](#chapter-6--pointer-stackheap-va-escape-analysis)
- [Chapter 7 — Type, method, interface](#chapter-7--type-method-interface)

---

## Chapter 1 — Go toolchain va Makefile

> Manba: `ch_1/text.txt`

### Asosiy buyruqlar

| Vazifa | Buyruq |
| --- | --- |
| Compile | `go build` |
| Code formatter | `go fmt` |
| Dependency manager | `go mod` |
| Test runner | `go test` |
| Kodda xatolarni skanerlash | `go vet` |

### `go mod init [name]`

Bu buyruqning asosiy maqsadi — joriy papkani Go module'ga aylantirish. Natijada bitta `go.mod`
fayli yaratiladi. Shu fayl Go toolchain'ga uchta narsani aytadi:

1. bu papka loyihaning root'i;
2. loyihaning nomi (masalan `hello_world`);
3. loyiha Go'ning qaysi versiyasi uchun yozilgani.

### `go build` nima qiladi?

1. **`go.mod` ni topadi** — module root va module nomini aniqlaydi.
2. **Paketlarni yig'adi** — joriy papkadagi `.go` fayllarni va ular import qilgan barcha paketlarni
   (standart kutubxona va dependency'lar) aniqlaydi.
3. **Kompilyatsiya qiladi** — har bir paket uchun:
   - *parse*: kod sintaksis daraxtiga (AST) aylantiriladi;
   - *type check*: tiplar tekshiriladi, ishlatilmagan import va o'zgaruvchilar xato beradi;
   - *SSA optimizatsiya*: inlining, escape analysis, keraksiz kodni olib tashlash;
   - mashina kodi generatsiyasi.
4. **Link qiladi** — barcha paketlar va Go runtime (garbage collector, goroutine scheduler) bitta
   faylga birlashtiriladi.

### `go fmt ./...`

Formatlashni bir xil ko'rinishga keltirish uchun.

### `go vet`

Bir turdagi buglar bor: kod sintaktik jihatdan to'g'ri, lekin ehtimol noto'g'ri ishlaydi.
`go vet` aynan shunday xatolarni aniqlash uchun ishlatiladi.

### Makefile

Umumiy qilib aytilsa — loyihaning **buyruqlar menyusi**.

```makefile
.DEFAULT_GOAL := build
```

`.DEFAULT_GOAL` default target'ni belgilaydi. Agar u yozilmasa, default target sifatida birinchi
yozilgan buyruq olinadi.

---

## Chapter 2 — Literal'lar, tiplar, type conversion va const

> Manba: `ch_2/text.txt`, `ch_2/type_conversion/text.txt`

### 4 ta asosiy literal

- integer literal
- floating-point literal
- rune literal
- string literal

### `integer` va `float` farqi

- **integer**: har bir bit sonning qiymatini ifodalashga xizmat qiladi.
- **float**: IEEE 754 standarti asosida ishlaydi va sonni 3 qismga ajratadi:
  1. **Sign** — son musbat yoki manfiy ekanini anglatadi;
  2. **Exponent** — sonning kattalik darajasini bildiradi;
  3. **Mantissa (significand)** — sonning muhim raqamlarini saqlaydi.

### String

- **immutable**;
- default qiymati: `""`;
- Go'da UTF-8 formatidagi baytlar ketma-ketligi, shu sababli istalgan Unicode belgini saqlashi mumkin.

### Rune

- `int32`ning alias'i va bitta Unicode kod nuqtasini ifodalaydi;
- bitta qo'shtirnoq bilan yoziladi: `'a'`.

### Type conversion

- hech bir tipni `bool`ga o'girib bo'lmaydi — na implicit, na explicit;
- Go avtomatik konvertatsiyaga yo'l qo'ymaydi;
- **implicit** = avtomatik konvertatsiya (Go'da bunday narsa yo'q);
- **explicit** = konvertatsiyani developer o'zi yozadi.

### Const

Go'dagi constant'lar — literal'larga nom berish usuli. Ular faqat compiler **compile time**'da
hisoblay oladigan qiymatlarni saqlashi mumkin.

**Compile time** — Go compiler kodni tekshirayotgan va executable dasturga aylantirayotgan vaqt.
**Runtime** — dastur ishga tushib, amallarni bajarayotgan vaqt.

### `const` va `var` farqi

| Xususiyat | `const` | `var` |
| --- | --- | --- |
| Compile time'da aniqlanadigan qiymat | Ha | Ha |
| Runtime qiymati | O'zgarmaydi | O'zgarishi mumkin |
| Qiymatini qayta tayinlash | Mumkin emas | Mumkin |
| Typed declaration | Ha | Ha |
| Untyped declaration | Ha | Yo'q |
| Funksiya natijasini saqlash | Mumkin emas | Mumkin |

Constant'lar compile time'da mavjud bo'ladi, variable'lar esa dastur xotirasidagi o'zgaruvchilar
sifatida ishlatiladi.

---

## Chapter 3 — Composite tiplar: array, slice, string, map, struct

### Array

> Manba: `ch_3/array/text.txt`

- aniq uzunlikka ega bo'lishi yoki initialize qilinishi shart;
- `[...]` shaklida yozilsa, qiymatlari berilishi shart;
- uzunligidan ko'p element qo'shilsa xatolik beradi;
- `==` orqali ikkita array'ning tengligini tekshirish mumkin.

### Slice

> Manba: `ch_3/slice/text.txt`

**slice vs array:** array qat'iy uzunlikka ega bo'lishi kerak; slice'da bu muammo yo'q va uni
xohlagancha kengaytirish mumkin.

Index bilan initialize qilish:

```go
[]int{1, 2, 3, 4: 10, 56} // => 1, 2, 3, 10, 56
```

**Slice comparable emas:** ikkita slice'ni `==` yoki `!=` bilan tenglikka tekshirib bo'lmaydi.
Taqqoslash kerak bo'lsa `slices.Equal()` ishlatiladi.

**Slice yaratish:**

```go
var x []int      // nil slice
var x = []int{}  // empty slice (x := []int{} bilan bir xil)

x := make([]int, 2)
fmt.Println(x, len(x), cap(x)) // [0 0] 2 2
```

`make` bilan yaratilganda `append` mavjud elementlardan **keyin** qo'shadi — ya'ni boshlang'ich 0
qiymatlar o'rnini bosmaydi.

**Capacity:**

> If you try to add additional values when the length equals the capacity, the `append` function uses
> the Go runtime to allocate a new backing array for the slice with a larger capacity.
> The values in the original backing array are copied to the new one, the new values are added to the
> end of the new backing array, and the slice is updated to refer to the new backing array. Finally,
> the updated slice is returned.

Slice `append` orqali o'sganda Go runtime yangi xotira ajratishi va eski ma'lumotni ko'chirishi uchun
vaqt ketadi. Eski xotira ham garbage collect qilinishi kerak. Shu sababli Go runtime capacity
tugaganda slice'ni odatda bittadan ko'proqqa oshiradi.

**Capacity'ni hisoblash (Go 1.18'dan):**

- hozirgi `cap` 256 dan kichik bo'lsa — 2x ga oshadi;
- 256 ga teng yoki undan katta bo'lsa — qo'shiladigan miqdor `(current_cap + 768) / 4`,
  jami capacity esa `current_cap + (current_cap + 768) / 4`.

Bu sekin-asta 25% o'sishga yaqinlashadi (capacity 512 bo'lsa 63% ga, 4096 bo'lsa faqat 30% ga o'sadi).

**`make`:** slice yaratishda `len` qancha berilsa, shuncha elementga initial value beriladi.

**`clear`:** `clear` funksiyasi slice'ning barcha qiymatlarini zero value'ga o'tkazadi.

**Slicing a slice:** slicing qilganda data copy qilinmaydi — shuning uchun o'zgarishlar bir-biriga
ta'sir qiladi (backing array va cap umumiy bo'ladi). Agar slicing'dan keyin cap oshsa (yangi backing
array ajratilsa), o'zgarishlar ta'sir qilmaydi.

```go
s[low:high:max] // max orqali capacity'ni cheklash mumkin
```

### String

> Manba: `ch_3/string/text.txt`

String'ni slice qilish faqat string bir baytli harflardan tashkil topgan bo'lsa xavfsiz; qolgan
hollarda (ko'p baytli Unicode belgilar) xatolikka olib kelishi mumkin.

### Map

> Manba: `ch_3/map/text.txt`

```go
totalWins := map[string]int{}
```

Bu — empty map literal, va u **nil map bilan bir xil emas**. Uzunligi 0, lekin empty map literal
tayinlangan map'ga o'qish ham, yozish ham mumkin.

- map uchun zero value — `nil`;
- map'lar comparable emas: `nil` ga tengligini tekshirish mumkin, lekin ikkita map'ning kalit va
  qiymatlari bir xilligini `==` / `!=` bilan tekshirib bo'lmaydi.

**Map yoki slice — qaysi birini tanlash?**

- ma'lumot ketma-ket qayta ishlanishi kerak bo'lsa yoki elementlar tartibi muhim bo'lsa — **slice**;
- qiymatlarni o'sib boruvchi integer emas, boshqa narsa (masalan nom) bo'yicha tashkil qilish kerak
  bo'lsa — **map**.

**Map'dan element o'chirish:** map `nil` bo'lsa yoki element mavjud bo'lmasa, hech narsa sodir
bo'lmaydi; `delete` hech narsa qaytarmaydi.

**Map'larni taqqoslash:** `maps.Equal` va `maps.EqualFunc` (`slices.Equal` va `slices.EqualFunc` kabi).

### Struct

> Manba: `ch_3/struct/text.txt` — hozircha bo'sh, qaydlar yozilmagan.

---

## Chapter 4 — Blok va shadowing

> Manba: `ch_4/shadow/text.txt`

**Shadowing:** shadowing variable — bu o'zini o'z ichiga olgan blokdagi o'zgaruvchi bilan bir xil
nomga ega o'zgaruvchi. Shadowing variable mavjud bo'lib turganda, shadow qilingan (tashqi)
o'zgaruvchiga murojaat qila olmaysiz.

---

## Chapter 5 — Funksiyalar va defer

> Manba: `ch_5/function/text.txt`, `ch_5/f/text.txt`, `ch_5/pass_value_type/text.txt`

### Funksiyalar

- string'dan int'ga: `strconv.Atoi`;
- funksiyadan qaytgan qiymatni `var` orqali olish mumkin, `const` orqali esa yo'q.

### Defer

`defer` **LIFO** (last in, first out) tartibida ishlaydi — oxirgi `defer` qilingan chaqiruv birinchi
bajariladi.

### Pass by value

> `ch_5/pass_value_type/text.txt` — hozircha bo'sh; kod `ch_5/pass_value_type/exercise.go` da.

---

## Chapter 6 — Pointer, stack/heap va escape analysis

### Pointer

> Manba: `ch_6/pointer/text.txt`

**Pointer ta'rifi:** pointer — bu boshqa bir o'zgaruvchining qiymati saqlangan xotira manzilini
o'zida ushlab turuvchi o'zgaruvchi.

**Address operatori:** `&` — qiymat oldiga qo'yilib, uning xotira manzilini qaytaradi.

**Map:** Go runtime tizimida `map` ichki struct'ga pointer sifatida amalga oshirilgan. Shuning uchun
map funksiyaga uzatilganda pointer'ning nusxasi o'tadi va map elementlarini o'zgartirish, qo'shish
yoki o'chirish chaqiruvchida ham ta'sir etadi.

**Slice:** slice uchta maydondan iborat kichik struktura:

1. boshlang'ich massivga pointer (`array`);
2. uzunlik (`len`);
3. sig'im (`cap`).

### Stack, Heap va Escape Analysis

> Manba: `ch_6/memory/text.txt`

**Stack (stek):** ketma-ket xotira bloki bo'lib, unda xotira ajratish va tozalash juda tez (stack
pointer'ini siljitish orqali) amalga oshadi. Hajmi kompilyatsiya vaqtidayoq ma'lum bo'lgan barcha
qiymatlar (primitivlar, struct'lar, massivlar va hatto pointer'larning o'zi) stack'da saqlanadi.

**Heap:** Garbage Collector tomonidan boshqariladigan dinamik xotira.

**Escape Analysis:** Go kompilyatori o'zgaruvchining qayerda saqlanishini tahlil qiladi. Agar
funksiya ichida yaratilgan mahalliy o'zgaruvchining pointer'i funksiyadan tashqariga qaytarilsa, u
stack'dan "qochadi" (*escapes to the heap*) va avtomatik ravishda heap'da saqlanadi. Shu sababli
Go'da lokal o'zgaruvchining pointer'ini qaytarish butunlay xavfsiz.

**Struct slice'lari:** xotirada ketma-ket joylashgani uchun protsessor keshi orqali juda tez
o'qiladi. Struct pointer'lari slice'i esa RAM bo'ylab tarqoq joylashadi va ularni o'qish taxminan
ikki barobar sekinroq kechadi.

---

## Chapter 7 — Type, method, interface

> Manba: `ch_7/types/text.txt` — hozircha bo'sh. Chapter boshlangan (`types`, `methods`,
> `interfaces` papkalari), lekin qaydlar hali yozilmagan.
