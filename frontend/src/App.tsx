import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import BlockGrid from './pages/BlockGrid';
import AddBlock from './pages/AddBlock';
import BlockDetails from './pages/BlockDetails';
import SimulateTx from './pages/SimulateTx';

function App() {
  return (
    <Router>
      <nav style={{ padding: '10px', background: '#eee' }}>
        <Link to="/" style={{ marginRight: '10px' }}>Grid</Link>
        <Link to="/add">Add Block</Link>
        <Link to="/details" style={{ marginLeft: '10px' }}>Block Details</Link>
        <Link to="/simulate" style={{ marginLeft: '10px' }}>Simulate Tx</Link>
      </nav>

      <Routes>
        <Route path="/" element={<BlockGrid />} />
        <Route path="/add" element={<AddBlock />} />
        <Route path="/details" element={<BlockDetails />} />
        <Route path="/simulate" element={<SimulateTx />} />
      </Routes>
    </Router>
  );
}

export default App;
