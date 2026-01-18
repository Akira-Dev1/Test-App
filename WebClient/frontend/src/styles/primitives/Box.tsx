import styled from 'styled-components';

interface BoxProps {
  padding?: number;
  gap?: number;
}

export const Box = styled.div<BoxProps>`
  display: flex;
  flex-direction: column;
  padding: ${({ padding }) => padding ?? 0}px;
  gap: ${({ gap }) => gap ?? 0}px;
`;