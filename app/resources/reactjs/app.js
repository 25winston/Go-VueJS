'use strict'

import './bootstrap';
import 'bootstrap';
import React from 'react';
import { render } from 'react-dom';


import Welcome from './components/Welcome';

render(
  <Welcome/>,
  document.getElementById('App'));