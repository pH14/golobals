# golobals

Keep config variables up-to-date and orderly with `golobals`.

Build a `Source` stack that defines a hierarchy for fetching named values. For instance, if your values come from ZooKeeper but you want a local YAML file to take priority if they exist, you might use:

    golobals := golobals.Create(YamlSource{"conf.yaml"}, ZookeeperSource{"zk.app.com"})

Then you can define your named values in your application either with struct tags:

    type AppConfig struct {
    	PostgresURL    LiveString `v:"app.postgres.url"`
    	RedisURL       LiveString `v:"app.redis.url"`
    	KeyczarHmacKey LiveString `v:"app.keyczar.hmac"`
    }
    
    appConfig := globals.Init(AppConfig{}).(AppConfig)

Or access them directly:

    redisDirect := golobals.Get("app.redis.url")

Now you can get your values with a simple `Get()` call, which will on-the-fly cycle through your Source stack until a value is found.

    redisUrl := AppConfig.RedisURL.Get()
    redisUrl := redisDirect.Get()

Sources can do whatever they need to do to fetch values, like HTTP requests or database lookups, so your application's config values can stay up-to-date each time they're accessed at runtime.

    type RemoteKVServer struct{}
    func (e *RemoteKVServer) Get(varName string) string {
      // Do fancy stuff here like an HTTP request
      return latestVal
    }
