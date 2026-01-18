import styled from 'styled-components';
import { theme } from '../theme';

export const Input = styled.input`
  background: ${theme.colors.surface};
  border: 1px solid ${theme.colors.border};
  padding: 10px 12px;
  border-radius: ${theme.radius.md};
  color: ${theme.colors.text};

  &:focus {
    outline: none;
    border-color: ${theme.colors.primary};
  }
`;