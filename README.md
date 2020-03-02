# oscsend
A simple tool for sending start/stop messages to [Max](https://cycling74.com)

## To Build
```
% git clone https://github.com/discursive-image/oscsend.git
% cd oscsend
% make
```

The executable will be available inside `bin/`.

## Usage
```
% bin/oscsend --help # for some usage help.
% bin/oscsend "path to an image" # to send a /max/play message with the path as payload.
% bin/oscsend # to send a /max/stop message.
```
