import styled from 'styled-components';

export const Stack = styled.div<{ gap?: number }>`
  display: flex;
  flex-direction: column;
  gap: ${({ gap }) => gap ?? 12}px;
`;