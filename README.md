# imptools

This is a suite of tools for reading and writing Imperium Galactica (1997) game data. There are currently 3 tools: `sav`, `unpac`, and `repac`.

## `sav`

`sav` enumerates the contents of an Imperium Galactica file. It is currently capable of reading the save header (rank, difficulty, money, date, etc), planets, buildings, technologies, fleets, and ships.
I have not completely identified every field on each data structure, and interpreting some of them requires data from the game binary that I am still working out, but it is reasonably complete.

### Usage

`sav` takes two positional arguments. It must be passed a path to an IG save file, and a set of flags identifying which structures to print.
(**h**eader, **p**lanets, **b**uildings, **t**echnologies, **f**leets, **s**hips)

```sh
sav /path/to/savefile hpbtfs
```

## `unpac`

`unpac` extracts the contents of a .PAC archive containing game data. It requres a path to an archive and an output path.
The output path must be empty; it will create a directory there to extract the files into.

```sh
unpac colony.pac colony
```

## `repac`

`repac` assembles a .PAC archive from the contents of a directory. It requies a path to an input directory and an output filename.
It will read files from the root of the input directory (ignoring other directories) and create the .PAC with the necessary headers.
Filenames must be 13 bytes or less, and there are limits on the number of files and total length including headers. (max uint16 and max uint32 respectively).

```sh
repac colony colony.pac
```
