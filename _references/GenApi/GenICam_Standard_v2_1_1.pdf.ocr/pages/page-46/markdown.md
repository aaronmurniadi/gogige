|  GENICAM |   | ![img-62.jpeg](img-62.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

- Int64 GetDisplayPrecision() – determines the precision to display the float number with
- IInteger *GetIntAlias() - returns a node with represents the same value in integer type
- IEnumeration *GetEnumAlias() - returns a node with represents the same value in enumeration type
- void ImposeMin( int64 ) – restricts the minimum
- void ImposeMax( int64 ) – restricts the maximum

#### 2.9.3 IString Interface

- string GetValue() – returns the value
- void SetValue( string ) – sets the value
- int64 GetMaxLength() – gets the maximum length of the string

#### 2.9.4 IEnumeration Interface

- int64 GetIntValue() – returns the index value corresponding to the enumeration value
- void SetIntValue( int64 ) – sets the index value corresponding to the enumeration value
- NodeList GetEntries() – returns a list of pointers to the EnumEntry nodes of the enumeration
- void GetSymbolics( StringList& ) – returns a list of valid enumeration values
- IEnumEntry *GetEntryByName( string ) – gets the EnumEntry corresponding to the symbolic
- IEnumEntry *GetEntry ( int64 ) – gets the EnumEntry corresponding to the integer value
- IEnumEntry *GetCurrentEntry ( int64 ) – gets the currently active EnumEntry

#### 2.9.5 ICommand Interface

- void Execute() – submits the command
- boolean IsDone() – returns true if the command has been executed; false as long as it still executes.

#### 2.9.6 IBoolean Interface

- boolean GetValue() – returns the value
- void SetValue( boolean ) – sets the value

#### 2.9.7 IRegister Interface

- void Get( uint8 *pBuffer, int64 Length ) – gets the register's content to a buffer
- void Set( uint8 *pBuffer, int64 Length ) – sets the register's content from a buffer
- int64 GetAddress() – gets the register's address
- int64 GetLength() – gets the register's length in bytes