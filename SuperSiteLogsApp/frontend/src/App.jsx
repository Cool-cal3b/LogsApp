import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import Home from './Home';
import Logs from './Logs';
import KeyPathSelect from './KeyPathRetriever'
import Log from './Log';

import '@fortawesome/fontawesome-free/css/all.min.css';

function App() {
    return (
        <Router>
            <nav id="MainNavBar">
                <ul>
                    <li><Link to="/">Home</Link></li>
                    <li><Link to="/logs">Logs</Link></li>
                    <li><Link to="/keySelect">Key Path Select</Link></li>
                </ul>
            </nav>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/logs" element={<Logs />} />
                <Route path="/keySelect" element={<KeyPathSelect />} />
                <Route path="/log/:id" element={<Log />} />
            </Routes>
        </Router>
    );
}

export default App;
