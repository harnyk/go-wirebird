import { createElement } from 'react';
import { render } from 'react-dom';

import { LoggerEvent } from '../generated/types';
import { MainScreen } from './src/screens/main.screen';

main();

function main() {
    initUI();
}

function initUI() {
    render(createElement(MainScreen), document.getElementById('root'));
}
