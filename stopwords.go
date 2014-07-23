package goose

import (
	"gopkg.in/fatih/set.v0"
	//"io/ioutil"
	"regexp"
	"strings"
)

var PUNCTUATION = regexp.MustCompile("[^\\p{Ll}\\p{Lu}\\p{Lt}\\p{Lo}\\p{Nd}\\p{Pc}\\s]")

type StopWords struct {
	cachedStopWords map[string]*set.Set
}

func NewStopwords() StopWords {
	cachedStopWords := make(map[string]*set.Set)
	for lang, stopwords := range sw {
		lines := strings.Split(stopwords, "\n")
		cachedStopWords[lang] = set.New()
		for _, line := range lines {
			line = strings.Trim(line, " ")
			line = strings.Trim(line, "\t")
			line = strings.Trim(line, "\r")
			cachedStopWords[lang].Add(line)
		}
	}
	return StopWords{
		cachedStopWords: cachedStopWords,
	}
}

/*
func NewStopwords(path string) StopWords {
	cachedStopWords := make(map[string]*set.Set)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err.Error())
	}
	for _, file := range files {
		name := strings.Replace(file.Name(), ".txt", "", -1)
		name = strings.Replace(name, "stopwords-", "", -1)
		name = strings.ToLower(name)

		stops := set.New()
		lines := ReadLinesOfFile(path + "/" + file.Name())
		for _, line := range lines {
			line = strings.Trim(line, " ")
			stops.Add(line)
		}
		cachedStopWords[name] = stops
	}

	return StopWords{
		cachedStopWords: cachedStopWords,
	}
}
*/

func (this *StopWords) removePunctuation(text string) string {
	return PUNCTUATION.ReplaceAllString(text, "")
}

func (this *StopWords) stopWordsCount(lang string, text string) wordStats {
	if text == "" {
		return wordStats{}
	}
	ws := wordStats{}
	stopWords := set.New()
	text = strings.ToLower(text)
	items := strings.Split(text, " ")
	stops := this.cachedStopWords[lang]
	count := 0
	if stops != nil {
		for _, item := range items {
			if stops.Has(item) {
				stopWords.Add(item)
				count++
			}
		}
	}

	ws.stopWordCount = stopWords.Size()
	ws.wordCount = len(items)
	ws.stopWords = stopWords

	return ws
}

var sw = map[string]string{
	"ar": `
فى
في
كل
لم
لن
له
من
هو
هي
قوة
كما
لها
منذ
وقد
ولا
نفسه
لقاء
مقابل
هناك
وقال
وكان
نهاية
وقالت
وكانت
للامم
فيه
كلم
لكن
وفي
وقف
ولم
ومن
وهو
وهي
يوم
فيها
منها
مليار
لوكالة
يكون
يمكن
مليون
حيث
اكد
الا
اما
امس
السابق
التى
التي
اكثر
ايار
ايضا
ثلاثة
الذاتي
الاخيرة
الثاني
الثانية
الذى
الذي
الان
امام
ايام
خلال
حوالى
الذين
الاول
الاولى
بين
ذلك
دون
حول
حين
الف
الى
انه
اول
ضمن
انها
جميع
الماضي
الوقت
المقبل
اليوم
ـ
ف
و
و6
قد
لا
ما
مع
مساء
هذا
واحد
واضاف
واضافت
فان
قبل
قال
كان
لدى
نحو
هذه
وان
واكد
كانت
واوضح
مايو
ب
ا
أ
،
عشر
عدد
عدة
عشرة
عدم
عام
عاما
عن
عند
عندما
على
عليه
عليها
زيارة
سنة
سنوات
تم
ضد
بعد
بعض
اعادة
اعلنت
بسبب
حتى
اذا
احد
اثر
برس
باسم
غدا
شخصا
صباح
اطار
اربعة
اخرى
بان
اجل
غير
بشكل
حاليا
بن
به
ثم
اف
ان
او
اي
بها
صفر
	`,
	"en": `
	a's
able
about
above
according
accordingly
across
actually
after
afterwards
again
against
ain't
all
allow
allows
almost
alone
along
already
also
although
always
am
among
amongst
an
and
another
any
anybody
anyhow
anyone
anything
anyway
anyways
anywhere
apart
appear
appreciate
appropriate
are
aren't
around
as
aside
ask
asking
associated
at
available
away
awfully
be
became
because
become
becomes
becoming
been
before
beforehand
behind
being
believe
below
beside
besides
best
better
between
beyond
both
brief
but
by
c
c'mon
c's
came
campaign
can
can't
cannot
cant
cause
causes
certain
certainly
changes
clearly
co
com
come
comes
concerning
consequently
consider
considering
contain
containing
contains
corresponding
could
couldn't
course
currently
definitely
described
despite
did
didn't
different
do
does
doesn't
doing
don't
done
down
downwards
during
each
edu
eight
either
else
elsewhere
enough
endorsed
entirely
especially
et
etc
even
ever
every
everybody
everyone
everything
everywhere
ex
exactly
example
except
far
few
fifth
first
financial
five
followed
following
follows
for
former
formerly
forth
four
from
further
furthermore
get
gets
getting
given
gives
go
goes
going
gone
got
gotten
greetings
had
hadn't
happens
hardly
has
hasn't
have
haven't
having
he
he's
hello
help
hence
her
here
here's
hereafter
hereby
herein
hereupon
hers
herself
hi
him
himself
his
hither
hopefully
how
howbeit
however
i'd
i'll
i'm
i've
if
ignored
immediate
in
inasmuch
inc
indeed
indicate
indicated
indicates
inner
insofar
instead
into
inward
is
isn't
it
it'd
it'll
it's
its
itself
just
keep
keeps
kept
know
knows
known
last
lately
later
latter
latterly
least
less
lest
let
let's
like
liked
likely
little
look
looking
looks
ltd
mainly
many
may
maybe
me
mean
meanwhile
merely
might
more
moreover
most
mostly
much
must
my
myself
name
namely
nd
near
nearly
necessary
need
needs
neither
never
nevertheless
new
next
nine
no
nobody
non
none
noone
nor
normally
not
nothing
novel
now
nowhere
obviously
of
off
often
oh
ok
okay
old
on
once
one
ones
only
onto
or
other
others
otherwise
ought
our
ours
ourselves
out
outside
over
overall
own
particular
particularly
per
perhaps
placed
please
plus
possible
presumably
probably
provides
quite
quote
quarterly
rather
really
reasonably
regarding
regardless
regards
relatively
respectively
right
said
same
saw
say
saying
says
second
secondly
see
seeing
seem
seemed
seeming
seems
seen
self
selves
sensible
sent
serious
seriously
seven
several
shall
she
should
shouldn't
since
six
so
some
somebody
somehow
someone
something
sometime
sometimes
somewhat
somewhere
soon
sorry
specified
specify
specifying
still
sub
such
sup
sure
t's
take
taken
tell
tends
than
thank
thanks
thanx
that
that's
thats
the
their
theirs
them
themselves
then
thence
there
there's
thereafter
thereby
therefore
therein
theres
thereupon
these
they
they'd
they'll
they're
they've
think
third
this
thorough
thoroughly
those
though
three
through
throughout
thru
thus
to
together
too
took
toward
towards
tried
tries
truly
try
trying
twice
two
under
unfortunately
unless
unlikely
until
unto
up
upon
us
use
used
useful
uses
using
usually
uucp
value
various
very
via
viz
vs
want
wants
was
wasn't
way
we
we'd
we'll
we're
we've
welcome
well
went
were
weren't
what
what's
whatever
when
whence
whenever
where
where's
whereafter
whereas
whereby
wherein
whereupon
wherever
whether
which
while
whither
who
who's
whoever
whole
whom
whose
why
will
willing
wish
with
within
without
won't
wonder
would
would
wouldn't
yes
yet
you
you'd
you'll
you're
you've
your
yours
yourself
yourselves
zero
official
sharply
criticized
`,
	"es": `
de
la
que
el
en
y
a
los
del
se
las
por
un
para
con
no
una
su
al
lo
como
más
pero
sus
le
ya
o
este
sí
porque
esta
entre
cuando
muy
sin
sobre
también
me
hasta
hay
donde
quien
desde
todo
nos
durante
todos
uno
les
ni
contra
otros
ese
eso
ante
ellos
e
esto
mí
antes
algunos
qué
unos
yo
otro
otras
otra
él
tanto
esa
estos
mucho
quienes
nada
muchos
cual
poco
ella
estar
estas
algunas
algo
nosotros
mi
mis
tú
te
ti
tu
tus
ellas
nosotras
vosotros
vosotras
os
mío
mía
míos
mías
tuyo
tuya
tuyos
tuyas
suyo
suya
suyos
suyas
nuestro
nuestra
nuestros
nuestras
vuestro
vuestra
vuestros
vuestras
esos
esas
estoy
estás
está
estamos
estáis
están
esté
estés
estemos
estéis
estén
estaré
estarás
estará
estaremos
estaréis
estarán
estaría
estarías
estaríamos
estaríais
estarían
estaba
estabas
estábamos
estabais
estaban
estuve
estuviste
estuvo
estuvimos
estuvisteis
estuvieron
estuviera
estuvieras
estuviéramos
estuvierais
estuvieran
estuviese
estuvieses
estuviésemos
estuvieseis
estuviesen
estando
estado
estada
estados
estadas
estad
he
has
ha
hemos
habéis
han
haya
hayas
hayamos
hayáis
hayan
habré
habrás
habrá
habremos
habréis
habrán
habría
habrías
habríamos
habríais
habrían
había
habías
habíamos
habíais
habían
hube
hubiste
hubo
hubimos
hubisteis
hubieron
hubiera
hubieras
hubiéramos
hubierais
hubieran
hubiese
hubieses
hubiésemos
hubieseis
hubiesen
habiendo
habido
habida
habidos
habidas

# forms of ser, to be (not including the infinitive):
soy
eres
es
somos
sois
son
sea
seas
seamos
seáis
sean
seré
serás
será
seremos
seréis
serán
sería
serías
seríamos
seríais
serían
era
eras
éramos
erais
eran
fui
fuiste
fue
fuimos
fuisteis
fueron
fuera
fueras
fuéramos
fuerais
fueran
fuese
fueses
fuésemos
fueseis
fuesen
siendo
sido
tengo
tienes
tiene
tenemos
tenéis
tienen
tenga
tengas
tengamos
tengáis
tengan
tendré
tendrás
tendrá
tendremos
tendréis
tendrán
tendría
tendrías
tendríamos
tendríais
tendrían
tenía
tenías
teníamos
teníais
tenían
tuve
tuviste
tuvo
tuvimos
tuvisteis
tuvieron
tuviera
tuvieras
tuviéramos
tuvierais
tuvieran
tuviese
tuvieses
tuviésemos
tuvieseis
tuviesen
teniendo
tenido
tenida
tenidos
tenidas
tened
`,
	"fr": `
	# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#-----------------------------------------------------------------------
# a couple of test stopwords to test that the words are really being
# configured from this file:
stopworda
stopwordb

#Standard english stop words taken from Lucene's StopAnalyzer
a
an
and
are
as
at
be
but
by
for
if
in
into
is
it
no
not
of
on
or
s
such
t
that
the
their
then
there
these
they
this
to
was
will
with
au
aux
avec
ce
ces
dans
de
des
du
elle
en
et
eux
il
je
la
le
leur
lui
ma
mais
me
même
mes
moi
mon
ne
nos
notre
nous
on
ou
par
pas
pour
qu
que
qui
sa
se
ses
son
sur
ta
te
tes
toi
ton
tu
un
une
vos
votre
vous
c
d
j
l
à
m
n
s
t
y
été
étée
étées
étés
étant
suis
es
est
sommes
êtes
sont
serai
seras
sera
serons
serez
seront
serais
serait
serions
seriez
seraient
étais
était
étions
étiez
étaient
fus
fut
fûmes
fûtes
furent
sois
soit
soyons
soyez
soient
fusse
fusses
fût
fussions
fussiez
fussent
ayant
eu
eue
eues
eus
ai
as
avons
avez
ont
aurai
auras
aura
aurons
aurez
auront
aurais
aurait
aurions
auriez
auraient
avais
avait
avions
aviez
avaient
eut
eûmes
eûtes
eurent
aie
aies
ait
ayons
ayez
aient
eusse
eusses
eût
eussions
eussiez
eussent
ceci
celà
cet
cette
ici
ils
les
leurs
quel
quels
quelle
quelles
sans
soi

`,
	"sv": `
#-----------------------------------------------------------------------
# translated
#-----------------------------------------------------------------------

kunna
om
ovan
enligt
i enlighet med detta
över
faktiskt
efter
efteråt
igen
mot
är inte
alla
tillåta
tillåter
nästan
ensam
längs
redan
också
även om
alltid
am
bland
bland
en
och
en annan
någon
någon
hur som helst
någon
något
ändå
ändå
var som helst
isär
visas
uppskatta
lämpligt
är
inte
runt
som
åt sidan
be
frågar
associerad
vid
tillgängliga
bort
väldigt
vara
blev
eftersom
bli
blir
blir
varit
innan
förhand
bakom
vara
tro
nedan
bredvid
förutom
bäst
bättre
mellan
bortom
både
kort
men
genom
c
c'mon
c: s
kom
kampanj
kan
kan inte
kan inte
cant
orsaka
orsaker
viss
säkerligen
förändringar
klart
co
com
komma
kommer
om
följaktligen
överväga
överväger
innehålla
innehållande
innehåller
motsvarande
kunde
kunde inte
kurs
närvarande
definitivt
beskrivits
trots
gjorde
inte
olika
göra
gör
inte
gör
inte
gjort
ned
nedåt
under
varje
edu
åtta
antingen
annars
någon annanstans
tillräckligt
godkändes
helt
speciellt
et
etc
även
någonsin
varje
alla
alla
allt
överallt
ex
exakt
exempel
utom
långt
få
femte
först
finansiella
fem
följt
efter
följer
för
fd
tidigare
framåt
fyra
från
ytterligare
dessutom
få
blir
få
given
ger
gå
går
gå
borta
fick
fått
hälsningar
hade
hade inte
händer
knappast
har
har inte
ha
har inte
med
han
han är
hallå
hjälpa
hence
henne
här
här finns
härefter
härmed
häri
härpå
hennes
själv
hej
honom
själv
hans
hit
förhoppningsvis
hur
howbeit
dock
jag skulle
jag ska
jag är
jag har
om
ignoreras
omedelbar
i
eftersom
inc
indeed
indikera
indikerade
indikerar
inre
mån
istället
in
inåt
är
är inte
den
det skulle
det ska
det är
dess
själv
bara
hålla
håller
hålls
vet
vet
känd
sista
nyligen
senare
senare
latterly
minst
mindre
lest
låt
låt oss
liknande
gillade
sannolikt
lite
ser
ser
ser
ltd
huvudsakligen
många
kan
kanske
mig
betyda
under tiden
endast
kanske
mer
dessutom
mest
mestadels
mycket
måste
min
själv
namn
nämligen
nd
nära
nästan
nödvändigt
behöver
behov
varken
aldrig
ändå
ny
nästa
nio
ingen
ingen
icke
ingen
ingen
eller
normalt
inte
ingenting
roman
nu
ingenstans
uppenbarligen
av
off
ofta
oh
ok
okay
gammal
på
en gång
ett
ettor
endast
på
eller
andra
andra
annars
borde
vår
vårt
oss
ut
utanför
över
övergripande
egen
särskilt
särskilt
per
kanske
placeras
vänligen
plus
möjligt
förmodligen
förmodligen
ger
ganska
citera
kvartalsvis
snarare
verkligen
rimligen
om
oavsett
gäller
relativt
respektive
höger
sa
samma
såg
säga
säger
säger
andra
det andra
se
ser
verkar
verkade
informationsproblem
verkar
sett
själv
själva
förnuftig
skickas
allvarlig
allvarligt
sju
flera
skall
hon
bör
bör inte
eftersom
sex
så
några
någon
på något sätt
någon
något
sometime
ibland
något
någonstans
snart
sorry
specificerade
ange
ange
fortfarande
sub
sådan
sup
säker
t s
ta
tas
berätta
tenderar
än
tacka
tack
thanx
att
det är
brinner
den
deras
deras
dem
själva
sedan
därifrån
där
det finns
därefter
därigenom
därför
däri
theres
därpå
dessa
de
de hade
de kommer
de är
de har
tror
tredje
detta
grundlig
grundligt
de
though
tre
genom
hela
thru
sålunda
till
tillsammans
alltför
tog
mot
mot
försökte
försöker
verkligt
försök
försöker
två gånger
två
enligt
tyvärr
såvida inte
osannolikt
tills
åt
upp
på
oss
använda
används
användbar
använder
användning
vanligtvis
uucp
värde
olika
mycket
via
viz
vs
vill
vill
var
var inte
sätt
vi
vi skulle
vi kommer
vi är
vi har
välkommen
väl
gick
var
var inte
vad
vad är
oavsett
när
varifrån
närhelst
där
var är
varefter
medan
varigenom
vari
varpå
varhelst
huruvida
som
medan
dit
som
vem är
vem
hela
vem
vars
varför
kommer
villig
önskar
med
inom
utan
kommer inte
undrar
skulle
skulle inte
ja
ännu
ni
du skulle
kommer du
du är
du har
din
själv
er
noll
tjänsteman
skarpt
kritiserade
`,
	"zh": `
的
一
不
在
人
有
是
为
以
于
上
他
而
后
之
来
及
了
因
下
可
到
由
这
与
也
此
但
并
个
其
已
无
小
我
们
起
最
再
今
去
好
只
又
或
很
亦
某
把
那
你
乃
它
吧
被
比
别
趁
当
从
到
得
打
凡
儿
尔
该
各
给
跟
和
何
还
即
几
既
看
据
距
靠
啦
了
另
么
每
们
嘛
拿
哪
那
您
凭
且
却
让
仍
啥
如
若
使
谁
虽
随
同
所
她
哇
嗡
往
哪
些
向
沿
哟
用
于
咱
则
怎
曾
至
致
着
诸
自
`,
"ru": `
а
е
и
ж
м
о
на
не
ни
об
но
он
мне
мои
мож
она
они
оно
мной
много
многочисленное
многочисленная
многочисленные
многочисленный
мною
мой
мог
могут
можно
может
можхо
мор
моя
моё
мочь
над
нее
оба
нам
нем
нами
ними
мимо
немного
одной
одного
менее
однажды
однако
меня
нему
меньше
ней
наверху
него
ниже
мало
надо
один
одиннадцать
одиннадцатый
назад
наиболее
недавно
миллионов
недалеко
между
низко
меля
нельзя
нибудь
непрерывно
наконец
никогда
никуда
нас
наш
нет
нею
неё
них
мира
наша
наше
наши
ничего
начала
нередко
несколько
обычно
опять
около
мы
ну
нх
от
отовсюду
особенно
нужно
очень
отсюда
в
во
вон
вниз
внизу
вокруг
вот
восемнадцать
восемнадцатый
восемь
восьмой
вверх
вам
вами
важное
важная
важные
важный
вдали
везде
ведь
вас
ваш
ваша
ваше
ваши
впрочем
весь
вдруг
вы
все
второй
всем
всеми
времени
время
всему
всего
всегда
всех
всею
всю
вся
всё
всюду
г
год
говорил
говорит
года
году
где
да
ее
за
из
ли
же
им
до
по
ими
под
иногда
довольно
именно
долго
позже
более
должно
пожалуйста
значит
иметь
больше
пока
ему
имя
пор
пора
потом
потому
после
почему
почти
посреди
ей
два
две
двенадцать
двенадцатый
двадцать
двадцатый
двух
его
дел
или
без
день
занят
занята
занято
заняты
действительно
давно
девятнадцать
девятнадцатый
девять
девятый
даже
алло
жизнь
далеко
близко
здесь
дальше
для
лет
зато
даром
первый
перед
затем
зачем
лишь
десять
десятый
ею
её
их
бы
еще
при
был
про
процентов
против
просто
бывает
бывь
если
люди
была
были
было
будем
будет
будете
будешь
прекрасно
буду
будь
будто
будут
ещё
пятнадцать
пятнадцатый
друго
другое
другой
другие
другая
других
есть
пять
быть
лучше
пятый
к
ком
конечно
кому
кого
когда
которой
которого
которая
которые
который
которых
кем
каждое
каждая
каждые
каждый
кажется
как
какой
какая
кто
кроме
куда
кругом
с
т
у
я
та
те
уж
со
то
том
снова
тому
совсем
того
тогда
тоже
собой
тобой
собою
тобою
сначала
только
уметь
тот
тою
хорошо
хотеть
хочешь
хоть
хотя
свое
свои
твой
своей
своего
своих
свою
твоя
твоё
раз
уже
сам
там
тем
чем
сама
сами
теми
само
рано
самом
самому
самой
самого
семнадцать
семнадцатый
самим
самими
самих
саму
семь
чему
раньше
сейчас
чего
сегодня
себе
тебе
сеаой
человек
разве
теперь
себя
тебя
седьмой
спасибо
слишком
так
такое
такой
такие
также
такая
сих
тех
чаще
четвертый
через
часто
шестой
шестнадцать
шестнадцатый
шесть
четыре
четырнадцать
четырнадцатый
сколько
сказал
сказала
сказать
ту
ты
три
эта
эти
что
это
чтоб
этом
этому
этой
этого
чтобы
этот
стал
туда
этим
этими
рядом
тринадцать
тринадцатый
этих
третий
тут
эту
суть
чуть
тысяч`,
}
