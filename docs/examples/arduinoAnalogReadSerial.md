# Arduino Analog Read Serial

This example was translated from [www.arduino.cc](https://www.arduino.cc/en/Tutorial/AnalogReadSerial)

    Li 0.0
    
    /*
        AnalogReadSerial
        
        Reads an analog input on pin 0, prints the result to the Serial Monitor.
        Graphical representation is available using Serial Plotter (Tools > Serial Plotter menu).
        Attach the center pin of a potentiometer to pin A0, and the outside pins to +5V and ground.
        
        This example code is in the public domain.
        
        http://www.arduino.cc/en/Tutorial/AnalogReadSerial
    */
    
    import:
        serial
        analog
        delay
    
    // the setup routine runs once when you press reset:
    export setup fn():
        // initialize serial communication at 9600 bits per second:
        serial.begin(9600)
    
    // the loop routine runs over and over again forever:
    export loop fn():
        // read the input on analog pin 0:
        sensorValue int = analog.read(analog.A0)
        
        // print out the value you read:
        serial.printLn(sensorValue)
        delay(1) // delay in between reads for stability
