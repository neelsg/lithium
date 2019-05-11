# Arduino Debounce

This example was translated from [www.arduino.cc](https://www.arduino.cc/en/Tutorial/Debounce)

    Li 0.0
    
    /*
        Debounce
        
        Each time the input pin goes from LOW to HIGH (e.g. because of a push-button
        press), the output pin is toggled from LOW to HIGH or HIGH to LOW. There's a
        minimum delay between toggles to debounce the circuit (i.e. to ignore noise).
        
        The circuit:
        - LED attached from pin 13 to ground
        - pushbutton attached from pin 2 to +5V
        - 10 kilohm resistor attached from pin 2 to ground
        
        - Note: On most Arduino boards, there is already an LED on the board connected
          to pin 13, so you don't need any extra components for this example.
        
        created 21 Nov 2006 by David A. Mellis
        modified 30 Aug 2011 by Limor Fried
        modified 28 Dec 2012 by Mike Walters
        modified 30 Aug 2016 by Arturo Guadalupi
        
        This example code is in the public domain.
        
        http://www.arduino.cc/en/Tutorial/Debounce
    */
    
    import:
        pin
        digital
        timer
    
    private:
        // constants won't change. They're used here to set pin numbers:
        buttonPin const = 2                // the number of the pushbutton pin
        ledPin const = 13                  // the number of the LED pin
        
        // Variables will change:
        ledState int = pin.HIGH            // the current state of the output pin
        buttonState int                    // the current reading from the input pin
        lastButtonState int = pin.LOW      // the previous reading from the input pin
        
        // the following variables are unsigned longs because the time, measured in
        // milliseconds, will quickly become a bigger number than can be stored in an int.
        lastDebounceTime int.UInt64 = 0    // the last time the output pin was toggled
        debounceDelay int.UInt64 = 50      // the debounce time; increase if the output flickers
    
    export setup fn():
        pin.mode(buttonPin, pin.INPUT)
        pin.mode(ledPin, pin.OUTPUT)
        
        // set initial LED state
        digital.write(ledPin, ledState)
    
    export loop fn():
        // read the state of the switch into a local variable:
        reading int = digital.read(buttonPin)
        
        // check to see if you just pressed the button
        // (i.e. the input went from LOW to HIGH), and you've waited long enough
        // since the last press to ignore any noise:
        
        // If the switch changed, due to noise or pressing:
        if reading != lastButtonState:
            // reset the debouncing timer
            lastDebounceTime = timer.millis
        
        if (timer.millis - lastDebounceTime) > debounceDelay:
            // whatever the reading is at, it's been there for longer than the debounce
            // delay, so take it as the actual current state:
            
            // if the button state has changed:
            if reading != buttonState:
                buttonState = reading
            
            // only toggle the LED if the new button state is HIGH
            if buttonState == pin.HIGH:
                ledState = !ledState
        
        // set the LED:
        digital.write(ledPin, ledState)
        
        // save the reading. Next time through the loop, it'll be the lastButtonState:
        lastButtonState = reading
