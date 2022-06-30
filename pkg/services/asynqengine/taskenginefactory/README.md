# TaskEngineFactory

This factory is responsible for running all the asynq mux servers.  
A webhooks subscription, which states the following;  

1. I am interested in the following
2. Send it here

is given its own asynq queue.  This allows us to not have a single subscription stop others from completing.  It alone will go into a retry cycle and finally its failed messages will be put into an archive queue.  

## Background Subscription Fetcher

There is a go routine that runs to update any changes in subscriptions.  Any change will cause a cascade shutdown of all engines and the restart will now be compliant to the new subscription set.  Its simply easier to restart the entire thing on this sort of change.  

So, if there are 10 subscriptions, there are going to be 10 asynq mux servers spun up to match to each subscriptions queue.  

