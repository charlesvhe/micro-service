using Newtonsoft.Json;
using Newtonsoft.Json.Serialization;

namespace Common
{
    public static class JsonExtensions
    {
        /// <summary>
        /// Convert the object to Json string
        /// </summary>
        /// <param name="obj">object</param>
        /// <param name="camelCase">using the camelCase style</param>
        /// <param name="includeNullValue">convert the null value of the object property</param>
        /// <returns></returns>
        public static string ToJsonString(this object obj, bool camelCase = true, bool includeNullValue = false)
        {
            var options = new JsonSerializerSettings()
            {
                NullValueHandling = includeNullValue
                    ? NullValueHandling.Include
                    : NullValueHandling.Ignore,
                ReferenceLoopHandling = ReferenceLoopHandling.Ignore
            };

            if (camelCase)
            {
                options.ContractResolver = new CamelCasePropertyNamesContractResolver();
            }

            return JsonConvert.SerializeObject(obj, options);
        }

        /// <summary>
        /// Deserialize Json object from the string
        /// </summary>
        /// <param name="jsonString">json string</param>
        /// <typeparam name="T">genericity of the object</typeparam>
        /// <returns></returns>
        public static T DeserializeJsonObject<T>(this string jsonString)
        {
            return JsonConvert.DeserializeObject<T>(jsonString);
        }
    }
}
