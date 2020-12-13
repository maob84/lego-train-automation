# lego-train-automation

Small project to automate my 9V Lego toy train. Currently this train is
controlled by +/-9V current with a haptic controller.

General idea is:

* Replace haptic controller with software interface:
  * LM298
  * Raspberry Pi
  * REST API
* Implement simple iOS app to control train via an old iPad mini

## Raspberry Pi

I just used an old (first generation) Raspberry Pi I found in my junk box. I downloaded an up-to-date Raspbian image (light version without desktop) and put it on an SD card. After adding a wifi USB stick I connected it to my local network.

## LM298

In my junk box I also found an LM298 module that can be controlled with PWM to provide variable DC current from -12 to +12V. In general I followed https://www.youtube.com/watch?v=2bganVdLg5Q&feature=youtu.be to connect the LM298 module to the GPIOs pin of the Raspberry Pi.