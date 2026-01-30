import { Card } from './Card';
import { Text } from '../primitives/Text';
import { Stack } from '../primitives/Stack';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const FeatureItem = ({ title, description }: any) => (
  <Card>
    <Stack gap={8}>
      <Text variant="h3">{title}</Text>
      <Text variant="muted">{description}</Text>
    </Stack>
  </Card>
);
