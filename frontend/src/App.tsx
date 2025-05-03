import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import GridPage from './pages/GridPage';
import AddBlockPage from './pages/AddBlockPage';
import BlockDetailPage from './pages/BlockDetailPage';
import SimulateTxPage from './pages/SimulateTxPage';
import DescriptionPage from './pages/DescriptionPage';
import HomePage from './pages/Home'; // if exists

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<GridPage />} />
        <Route path="/add" element={<AddBlockPage />} />
        <Route path="/block/:id" element={<BlockDetailPage />} />
        <Route path="/simulate" element={<SimulateTxPage />} />
        <Route path="/description" element={<DescriptionPage />} />
        <Route path="/home" element={<HomePage />} />
      </Routes>
    </Router>
  );
};

export default App;
