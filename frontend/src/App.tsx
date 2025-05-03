// src/App.tsx
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import GridPage from './pages/Grid';
import AddBlockPage from './pages/AddBlock';
import BlockDetailPage from './pages/BlockDetail';
import SimulateTxPage from './pages/Transactions';
import DescriptionPage from './pages/Description';
import HomePage from './pages/Home';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/grid" element={<GridPage />} />
        <Route path="/add" element={<AddBlockPage />} />
        <Route path="/block/:id" element={<BlockDetailPage />} />
        <Route path="/simulate" element={<SimulateTxPage />} />
        <Route path="/description" element={<DescriptionPage />} />
      </Routes>
    </Router>
  );
};

export default App;
