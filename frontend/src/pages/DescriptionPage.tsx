import React, { useEffect } from 'react';

const DescriptionPage: React.FC = () => {
  useEffect(() => {
    // Initialize particles.js and Three.js here
    particlesJS('particles-js', {
      particles: {
        number: { value: 80, density: { enable: true, value_area: 800 } },
        color: { value: '#0ff' },
        shape: { type: 'circle' },
        opacity: { value: 0.3 },
        size: { value: 4 },
        line_linked: { enable: true, distance: 130, color: '#0ff', opacity: 0.4, width: 1 },
        move: { enable: true, speed: 2 }
      },
      interactivity: {
        detect_on: 'canvas',
        events: {
          onhover: { enable: true, mode: 'repulse' },
          onclick: { enable: true, mode: 'push' },
          resize: true
        },
        modes: {
          repulse: { distance: 100, duration: 0.4 },
          push: { particles_nb: 4 }
        }
      },
      retina_detect: true
    });

    // Three.js setup can be done here
  }, []);

  return (
    <div>
      <h1>Description of AECH and Block Plain</h1>
      <div id="particles-js" style={{ width: '100%', height: '100%' }}></div>
    </div>
  );
};

export default DescriptionPage;
