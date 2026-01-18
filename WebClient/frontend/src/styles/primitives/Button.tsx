import styled from 'styled-components';
import { theme } from '../theme';

export const Button = styled.button`
  background: ${theme.colors.primary};
  color: white;
  border: none;
  padding: 10px 16px;
  border-radius: ${theme.radius.md};
  cursor: pointer;

  &:hover {
    background: ${theme.colors.primaryHover};
  }
`;