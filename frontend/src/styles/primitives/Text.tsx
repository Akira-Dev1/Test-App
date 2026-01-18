import styled, { css } from 'styled-components';
import { theme } from '../theme';

type Variant = 'h1' | 'h2' | 'h3' | 'body' | 'muted';

export const Text = styled.p<{ variant?: Variant }>`
  margin: 0;
  color: ${theme.colors.text};

  ${({ variant }) =>
    variant === 'h1' &&
    css`
      font-size: 28px;
      font-weight: 600;
    `}

  ${({ variant }) =>
    variant === 'h2' &&
    css`
      font-size: 22px;
      font-weight: 600;
    `}

  ${({ variant }) =>
    variant === 'h3' &&
    css`
      font-size: 18px;
      font-weight: 500;
    `}

  ${({ variant }) =>
    variant === 'muted' &&
    css`
      color: ${theme.colors.textMuted};
    `}
`;