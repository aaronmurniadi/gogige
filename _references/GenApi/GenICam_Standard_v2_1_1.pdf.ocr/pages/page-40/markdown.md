|  GEN<i>CAM |   | ![img-55.jpeg](img-55.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

When embedding formulas in XML files the problem arises that the characters <, >, and & cannot be used directly because they are part of the XML syntax. There are two possible solutions for that problem.

First you can escape these letters as follows:

< becomes &lt; (lt = thess than)
> becomes &gt; (gt = greater than)
& becomes &amp; (amp = ampersand)

As a result the formula  \( (x>0) \)  &&  \( (x<10) \)  becomes

<formula>(x &gt; 0) &amp;&amp; (x &lt; 10)</formula>

Alternatively you can declare the whole formula as non-XML-text by bracketing it with <![CDATA[ and ]]>. The formula then becomes:

<formula><![CDATA[ (x>0) && (x<10) ]]>/formula>

The SwissKnife syntax has some extensions: You can use named constants using the Constant entry and named sub expressions using the Expression entry as shown in the following example. The sub expressions may not refer to other sub expressions.

<SwissKnife Name="Result">
    ValueX
    ValueY
    2.0
    Expression Name="TwoX">2.0*X
    <Formula> TwoX * Y + Two </Formula>
</SwissKnife>

In Addition you can access the minimum, maximum, and increment of a node by using the variable name extensions .Min, .Max, .Inc, and – for completeness – also .Value. In addition .Entry.Name is allowed which accesses the integer value of an EnumEntry described by Name. As an example the SwissKnife to find the middle of the [min, max] range would look like this:

<SwissKnife Name="MidRange">
    Gain
    Gain
    Gain
    (Gain.Max - Gain.Min) / 2 </Formula>
</SwissKnife>

In contrast to the SwissKnife the Converter works bi-directionally. It implements an IFloat interface, looks a bit like the SwissKnife but contains an additional  element which can point to an IInteger or IFloat interface. It has two formulas: the <FormulaFrom> describes how to convert a value from the <pValue> node to the use domain and the <FormulaTo> describes how to convert a user value to the <pValue> domain. The <Slope> entry indicates if the formula is monotonously Increasing or Decreasing, if it is Varying (in this case the full number range is used), or if the slope is determined in an Automatic way by</Slope></pValue></pValue></pValue></pValue></pValue>