# Arduino Analog Read Serial

This example was translated from [www.arduino.cc](https://www.arduino.cc/en/Tutorial/AnalogReadSerial)

    # AnalogReadSerial
    #
    # Reads an analog input on pin 0, prints the result to the Serial Monitor.
    # Graphical representation is available using Serial Plotter (Tools > Serial Plotter menu).
    # Attach the center pin of a potentiometer to pin A0, and the outside pins to +5V and ground.
    #
    # This example code is in the public domain.
    #
    # http://www.arduino.cc/en/Tutorial/AnalogReadSerial
    
    setup func:
        # the setup routine runs once when you press reset
        
        # initialize serial communication at 9600 bits per second:
        serial.begin(9600)
    
    loop func:
        # the loop routine runs over and over again forever
        
        # read the input on analog pin 0:
        sensorValue int = analog.read(0)
        
        # print out the value you read:
        serial.printLn(sensorValue)
        timer.delay(1) # delay in between reads for stability
