|  GENICAM |   | ![img-54.jpeg](img-54.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

<IntSwissKnife Name="XTimesY">
    XValue
    YValue
    X*Y
    
</IntSwissKnife>

<Integer Name="XValue">
    <Value>42</Value>
</Integer>

<Integer Name="YValue">
    <Value>12</Value>
</Integer>

The \( <Formula> \) element contains a mathematical formula that can refer to variables defined by \( <pVariable> \) elements which point to an Integer node and have a Name attribute that defines the name of the variable inside the formula. The variable name must be upper case.

The Swiss knife used in the reference implementation is quite powerful. However, to simplify the task for people wanting to do their own implementation, the standard only allows a restricted set of mathematical operations. The following operations are supported by the standard:

( ) brackets
+ - * / addition, subtraction, multiplication, division
% remainder
** power
& | ^ ~ bitwise and / or / xor / not
&lt; &gt; = &gt; &lt; &lt; = &gt;= logical relations not equal / equal / greater / less / less of equal / greater or equal
&& || logical and / or
&lt;&lt; &gt;&gt; shift left, shift right

Conditional operator:

<condition> ? <true expr.=""> : <false expr.=""></false></true></condition>

Functions:

SGN, NEG,

Functions present only with the SwissKnife but not with the IntSwissKnife:

ATAN, COS, SIN, TAN, ABS, EXP, LN, LG, SQRT, TRUNC, FLOOR, CEIL, ROUND(x, precision=0), ASIN, ACOS, SGN, NEG, E, PI