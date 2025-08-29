#!/bin/bash
until awslocal health status | grep "running"; do
    sleep 1
done

echo "Verifying SES email identity..."
awslocal ses verify-email-identity --email-address test@example.com
echo "SES setup complete!"