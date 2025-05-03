import { useEffect } from 'react';
import 'particles.js';

declare global {
  interface Window {
    particlesJS: any;
  }
}

const DescriptionPage = () => {
  useEffect(() => {
    window.particlesJS('particles-js', {
      particles: {
        number: { value: 80, density: { enable: true, value_area: 800 } },
        color: { value: '#0ff' },
        shape: { type: 'circle' },
        opacity: { value: 0.5 },
        size: { value: 3 },
        line_linked: {
          enable: true,
          distance: 150,
          color: '#0ff',
          opacity: 0.4,
          width: 1,
        },
        move: { enable: true, speed: 2 },
      },
      interactivity: {
        detect_on: 'canvas',
        events: {
          onhover: { enable: true, mode: 'repulse' },
          onclick: { enable: true, mode: 'push' },
        },
      },
      retina_detect: true,
    });
  }, []);

  return <div id="particles-js" style={{ width: '100%', height: '100vh' }} />;
};

export default DescriptionPage;
