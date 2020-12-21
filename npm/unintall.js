function fn() {
    try {
        const binary = require('./bin');
        return binary();
    } catch (err) {
        // swallowing error because this script also runs as a preinstall hook and on the first installation
        // it wouldn't have any dependencies which would always result in an error
    }
}

const binary = fn();
if (binary) {
    binary.uninstall();
}
