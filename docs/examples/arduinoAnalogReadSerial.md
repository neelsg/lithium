# Arduino Analog Read Serial

This example was translated from [www.arduino.cc](https://www.arduino.cc/en/Tutorial/AnalogReadSerial)

    Li 0.0
    
    import:
        serial
        analog
        delay
    
    export setup fn():
        // initialize serial communication at 9600 bits per second:
        serial.begin(9600)
    
    export loop fn():
        // read the input on analog pin 0:
        sensorValue int = analog.read(analog.A0)
        
        // print out the value you read:
        serial.printLn(sensorValue)
        delay(1) // delay in between reads for stability
