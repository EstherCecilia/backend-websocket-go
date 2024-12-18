import { render, screen } from '@testing-library/react';
import App from './App';

test('renders text', () => {
  render(<App />);
  const linkElement = screen.getByText(/Notificação de Atualizações/i);
  expect(linkElement).toBeInTheDocument();
});
