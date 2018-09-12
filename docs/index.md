
## thinkpol

thinkpol controls [joycamp](https://github.com/abhishekkr/joycamp)

* locally go-routine execution of joycamp with given config

* schedule jobs of joycamp with job-config using `local`, `nomad`, `kubernetes`

> * local: Allows local exec of joycamp's Proc, for this to scale thinkpol itself need to be scaled behind load-balancer. But this wouldn't be much effective as LB might not keep track of active Procs while delegating new ones.
>
> * nomad: Delegate Nomad Job:Task to carry out Joycamp and thus scales to whatever infra nomad cluster can.
>
> * kubernetes: Similar to Nomad, handles Kubernetes' Jobs handling Joycamp Proc.

---

