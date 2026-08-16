|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

To provide a generic API on top of the File Access Controls, a FileAccessAdapter is defined in the GenApi. The Adapter provides two iostream interfaces to the device files:

- IDevFileStream Read from the device
- ODevFileStream Write to the device

The File Protocol Adapter is responsible for the mapping of the (I/O) DevFileStreamBuf actions Open, Close, UnderFlow, Overflow on File Access Controls

### Example Code for the streaminterface:

//GenApi::INodeMap * pInterface
ODevFileStream usersetWrite;
usersetWrite.open(pInterface, "UserSet1");
if( ! usersetWrite.fail() ) {
    usersetWrite << "Hello World\n";
}
usersetWrite.close();

IDevFileStream usersetRead;
usersetRead.open(pInterface, "UserSet1");
if( ! usersetRead.fail() ) {
    cout << usersetRead.rdbuf();
}
usersetRead.close();