import { Card } from './Card';
import { Stack } from '../primitives/Stack';
import { Text } from '../primitives/Text';

export const StatCard = ({ label, value }: { label: string; value: string }) => (
  <Card>
    <Stack gap={6}>
      <Text variant="h2">{value}</Text>
      <Text variant="muted">{label}</Text>
    </Stack>
  </Card>
);
