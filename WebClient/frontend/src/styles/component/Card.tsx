import styled from 'styled-components';
import { theme } from '../theme';

export const Card = styled.section`
  background: ${theme.colors.surface};
  border-radius: ${theme.radius.lg};
  padding: 24px;
  box-shadow: ${theme.shadows.card};
`;