# Rechner-Kata

Schreibe einen Rechner, der mathematische Ausdrücke liest, die
Fließkommazahlen, +, -, * und / sowie Klammern enthalten.

Berechne das Ergebnis, indem die "üblichen" mathematischen Regeln angewandt werden:

* Operatoren werden in den String-Inputs mit _Infix-Notation_ verwendet: `2 + 3`
* Multiplikation/Division haben eine höhere Priorität als Addition/Subtraktion
* Klammern haben eine höhere Priorität als Operatoren
* Zahlen können eine Mischung aus Ganzzahlen und Fließkommazahlen enthalten
* Die Berechnung sollte mathematischen Regeln folgen und `r / 0` (für jede Zahl r) ist ein Fehler und nicht
  etwas wie `NaN` oder `-Inf`

Der Rechner sollte

* das Ergebnis eines einzelnen mathematischen Ausdrucks berechnen (egal wie lang)
* das Ergebnis in einem menschenlesbaren Stil ausdrucken
* jeden Fehler melden und die Berechnung ablehnen (Stil der Fehlermeldung spielt keine Rolle)
* beliebig lange Ausdrücke gut verarbeiten (d.h. in Bezug auf Speicherverbrauch sowie Laufzeit)

