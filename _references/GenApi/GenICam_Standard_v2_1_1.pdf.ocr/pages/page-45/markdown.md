|  GEN<ICAM |   | ![img-61.jpeg](img-61.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

### 2.9 Available Interfaces

This section uses a pseudo code notation to list the most important interfaces as introduced in section 2.3. An actual implementation can have more methods per interface, e.g., in parallel to a SetValue(value) method, an operator=(value) method might be implemented that maps directly to the SetValue() method. Also, the actual variable types may differ, e.g., for the pseudo code type string, the actual implementation might be CString, std::string, or something else.

A more thorough explanation is found in section 2.8.

#### 2.9.1 Integer Interface

- int64 GetValue() – returns the value
- void SetValue(int64) – sets the value
- int64 GetMin() – returns the minimum
- int64 GetMax() – returns the maximum
- EIncMode GetIncMode() – returns the type of increment.
- int64 GetInc() – returns the increment if GetIncMode returns fixedIncrement
- gcint64_autovector GetListOfValidValue() returns a list of valid values if GetIncMode returns listIncrement
- ERepresentation GetRepresentation() – returns the representation as an enumeration
- string GetUnit() – returns the unit
- void ImposeMin(int64) – restricts the minimum
- void ImposeMax(int64) – restricts the maximum
- IFloat *GetFloatAlias() – returns a node with represents the same value in float type

#### 2.9.2 IFloat Interface

- double GetValue() – returns the value
- void SetValue(double) – sets the value
- double GetMin() – returns the minimum
- double GetMax() – returns the maximum
- EIncMode GetIncMode() - returns the type of increment
- int64 GetInc() – returns the increment if GetIncMode returns fixedIncrement
- gcdouble_autovector GetListOfValidValue() returns a list of valid value if GetIncMode returns listIncrement
- ERepresentation GetRepresentation() – returns the representation as an enumeration
- string GetUnit() – returns the unit
- EDisplayNotation GetDisplayNotation() – determines how to display the float number